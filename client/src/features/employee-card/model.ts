import { createEffect, createEvent, createStore, sample } from 'effector';
import { IEmployee, IEmployeeResponse } from '../../shared/types';

export interface QueryParams {
  firstName?: string;
  page: number;
  limit: number;
  position?: string[];
  department?: string[];
  subdivision?: string[];
  role?: string[];
  project?: string[];
  city?: string[];
}

export const $employeesStore = createStore<IEmployee[] | null>([]);
export const $activeFiltersStore = createStore<QueryParams>({ page: 1, limit: 10 });

export const getEmployeesFx = createEffect(async () => {
  const params = $activeFiltersStore.getState();
  const url = new URL('http://localhost:8080/api/employees');
  url.searchParams.append('page', params.page.toString());
  url.searchParams.append('limit', params.limit.toString());

  if (params.firstName) {
    url.searchParams.append('first_name_search', params.firstName);
  }

  const filterKeys = ['position', 'department', 'subdivision', 'role', 'project', 'city'] as const;
  filterKeys.forEach((key) => {
    if (params[key]) {
      params[key]!.forEach((value) => {
        url.searchParams.append(key, value);
      });
    }
  });

  const result = await fetch(url.toString(), {
    method: 'GET',
    credentials: 'include',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  const data = (await result.json()) as IEmployeeResponse;

  return data.employees;
});

export const updateFilter = createEvent<{ filter: keyof QueryParams; value: string }>();

$activeFiltersStore.on(updateFilter, (state, { filter, value }) => {
  const newState = { ...state };
  if (newState[filter]) {
    if (Array.isArray(newState[filter])) {
      const filterArray = newState[filter] as string[];
      if (filterArray.includes(value)) {
        newState[filter] = filterArray.filter((v) => v !== value);
      } else {
        newState[filter] = [...filterArray, value];
      }
    } else {
      newState[filter] = [value];
    }
  } else {
    newState[filter] = [value];
  }
  return newState;
});

sample({
  clock: getEmployeesFx.doneData,
  target: $employeesStore,
});
