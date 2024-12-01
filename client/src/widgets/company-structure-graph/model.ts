import { createEffect, createStore } from 'effector';
import { IEmployeeTree, IEmployeeTreeEl } from '../../shared/types';

// Мокированные данные
const mockData: IEmployeeTree = {
  id: '11',
  first_name: 'Петр',
  middle_name: 'Анатольевич',
  last_name: 'Сидоров',
  position: 'Аналитик',
  department: 'Департамент клиентской поддержки',
  subdivision: 'Отдел поддержки',
  role: 'Сотрудник',
  project: 'Проект Б',
  city: 'Санкт-Петербург',
  hire_date: '0001-01-01T00:00:00Z',
  subordinates: [
    {
      id: '166',
      first_name: 'Татьяна',
      middle_name: 'Юрьевна',
      last_name: 'Васильева',
      position: 'Аналитик',
      department: 'Департамент информационных технологий',
      subdivision: 'Отдел разработки',
      role: 'Сотрудник',
      project: 'Проект А',
      city: 'Москва',
      hire_date: '0001-01-01T00:00:00Z',
    },
    {
      id: '176',
      first_name: 'Евгений',
      middle_name: 'Артурович',
      last_name: 'Смирнов',
      position: 'Стажёр',
      department: 'Департамент информационных технологий',
      subdivision: 'Отдел разработки',
      role: 'Гостевой доступ',
      project: 'Проект А',
      city: 'Москва',
      hire_date: '0001-01-01T00:00:00Z',
    },
  ],
  colleagues: [
    {
      id: '10',
      first_name: 'Ольга',
      middle_name: 'Игоревна',
      last_name: 'Лебедева',
      position: 'Технический специалист',
      department: 'Департамент продаж',
      subdivision: 'Отдел продаж',
      role: 'Гостевой доступ',
      project: 'Проект А',
      city: 'Москва',
      hire_date: '0001-01-01T00:00:00Z',
    },
    {
      id: '12',
      first_name: 'Елена',
      middle_name: 'Викторовна',
      last_name: 'Григорьева',
      position: 'Бухгалтер',
      department: 'Маркетинговый департамент',
      subdivision: 'Маркетинговый отдел',
      role: 'Менеджер',
      project: 'Проект В',
      city: 'Новосибирск',
      hire_date: '0001-01-01T00:00:00Z',
    },
    {
      id: '162',
      first_name: 'Ольга',
      middle_name: 'Николаевна',
      last_name: 'Лебедева',
      position: 'Технический специалист',
      department: 'Департамент управления персоналом',
      subdivision: 'Отдел кадров',
      role: 'Менеджер',
      project: 'Проект Б',
      city: 'Санкт-Петербург',
      hire_date: '0001-01-01T00:00:00Z',
    },
    {
      id: '173',
      first_name: 'Марина',
      middle_name: 'Игоревна',
      last_name: 'Петрова',
      position: 'Специалист по продажам',
      department: 'Департамент информационных технологий',
      subdivision: 'Отдел разработки',
      role: 'Гостевой доступ',
      project: 'Проект Б',
      city: 'Москва',
      hire_date: '0001-01-01T00:00:00Z',
    },
  ],
  managers: [
    {
      id: '8',
      first_name: 'Мария',
      middle_name: 'Александровна',
      last_name: 'Смирнова',
      position: 'Аналитик',
      department: 'Департамент управления персоналом',
      subdivision: 'Отдел кадров',
      role: 'Менеджер',
      project: 'Проект Б',
      city: 'Санкт-Петербург',
      hire_date: '0001-01-01T00:00:00Z',
    },
  ],
};

export const getEmployeeTreeFx = createEffect(async (id: string) => {
  try {
    const response = await fetch(`http://localhost:8080/api/employees/tree/${id}`, {
      method: 'GET',
      credentials: 'include',
      mode: 'cors',
      headers: {
        'Content-Type': 'application/json',
      },
    });
    if (!response.ok) {
      return mockData;
    }
    return (await response.json()) as IEmployeeTree;
  } catch (e) {
    return mockData;
  }
});
export const $employeeTreeStore = createStore<IEmployeeTree | null>(null)
  .on(getEmployeeTreeFx.doneData, (_, employeeTree) => employeeTree)
  .reset(getEmployeeTreeFx.fail);

export const transformEmployeeTreeToGraph = (employeeTree: IEmployeeTree) => {
  const nodes = [];
  const edges = [];

  const addNode = (employee: IEmployeeTreeEl, position: { x: number; y: number }) => {
    nodes.push({
      id: employee.id.toString(),
      data: { label: `${employee.first_name} ${employee.last_name}` },
      position,
    });
  };

  const addEdge = (sourceId: string, targetId: string) => {
    edges.push({
      id: `e${sourceId}${targetId}`,
      source: sourceId,
      target: targetId,
      animated: true,
    });
  };

  const traverseTree = (employee: IEmployeeTree, position: { x: number; y: number }, level: number) => {
    addNode(employee, position);

    let yOffset = 100;
    let xOffset = 200;

    if (employee.colleagues) {
      employee.colleagues.forEach((colleague, index) => {
        const newPosition = { x: position.x + xOffset * (index + 1), y: position.y };
        addNode(colleague, newPosition);
        addEdge(employee.id.toString(), colleague.id.toString());
      });
    }

    if (employee.subordinates) {
      employee.subordinates.forEach((subordinate, index) => {
        const newPosition = { x: position.x + xOffset * (index + 1), y: position.y + yOffset };
        addNode(subordinate, newPosition);
        addEdge(employee.id.toString(), subordinate.id.toString());
      });
    }
  };

  const addManagers = (managers: IEmployeeTreeEl[], position: { x: number; y: number }) => {
    managers.forEach((manager, index) => {
      const newPosition = { x: position.x, y: position.y - 200 * (index + 1) };
      addNode(manager, newPosition);
      addEdge(manager.id.toString(), employeeTree.id.toString());
    });
  };

  if (employeeTree.managers) {
    addManagers(employeeTree.managers, { x: 0, y: 0 });
  }

  traverseTree(employeeTree, { x: 0, y: 100 }, 0);

  return { nodes, edges };
};
