import { createEffect, createStore, sample } from 'effector';
import { IFilter } from '../../shared/types';

export const $filtersStore = createStore<IFilter[]>([]);

export const getFiltersFx = createEffect(async () => {
  const response = await fetch('http://localhost:8080/api/employees/filters', {
    method: 'GET',
    credentials: 'include',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error('Failed to fetch filters');
  }

  return (await response.json()) as IFilter[];
});

$filtersStore.on(getFiltersFx.doneData, (_, filters) => filters);

sample({
  clock: getFiltersFx.doneData,
  target: $filtersStore,
});
