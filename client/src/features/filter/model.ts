import { createEffect, createStore, sample } from 'effector';
import { IFilter } from '../../shared/types';

export const $filtersStore = createStore<IFilter[]>([]);

export const getFiltersFx = createEffect(async () => {
  //   const result = await fetch('http://localhost:8080/api/employees/filters');

  //   return (await result.json()) as IFilter[];

  return [
    {
      filter: 'position',
      name: 'Должности',
      values: ['Аналитик', 'Бухгалтер', 'Менеджер', 'Специалист по продажам', 'Стажёр', 'Технический специалист'],
    },
    {
      filter: 'department',
      name: 'Департаменты',
      values: [
        'Департамент информационных технологий',
        'Департамент клиентской поддержки',
        'Департамент продаж',
        'Департамент управления персоналом',
        'Маркетинговый департамент',
        'Финансовый департамент',
      ],
    },
    {
      filter: 'subdivision',
      name: 'Подразделения',
      values: [
        'Маркетинговый отдел',
        'Отдел кадров',
        'Отдел поддержки',
        'Отдел продаж',
        'Отдел разработки',
        'Финансовый отдел',
      ],
    },
    {
      filter: 'role',
      name: 'Роли',
      values: ['Администратор', 'Гостевой доступ', 'Менеджер', 'Сотрудник'],
    },
    {
      filter: 'project',
      name: 'Проекты',
      values: ['Проект А', 'Проект Б', 'Проект В'],
    },
    {
      filter: 'city',
      name: 'Города',
      values: ['Москва', 'Новосибирск', 'Санкт-Петербург'],
    },
  ];
});

sample({
  clock: getFiltersFx.doneData,
  target: $filtersStore,
});
