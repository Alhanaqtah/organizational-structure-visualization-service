import { createEffect, createStore, sample } from 'effector';
import { IFilter } from '../../shared/types';

export const $filtersStore = createStore<IFilter[]>([]);

export const getFiltersFx = createEffect(async () => {
  const result = await fetch('http://localhost:8080/api/employees/filters');

  return (await result.json()) as IFilter[];
});

sample({
  clock: getFiltersFx.doneData,
  target: $filtersStore,
});
