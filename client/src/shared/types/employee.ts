export interface IEmployee {
  id: string;
  first_name: string;
  middle_name: string;
  last_name: string;
  position: string;
  department: string;
  subdivision: string;
  role: string;
  project: string;
  city: string;
  hire_date: string;
}

export interface IEmployeeResponse {
  total: number;
  page: number;
  limit: number;
  employees: IEmployee[];
}

export interface IEmployeeTreeEl {
  id: string;
  first_name: string;
  middle_name: string;
  last_name: string;
  position: string;
  department: string;
  subdivision: string;
  role: string;
  project: string;
  city: string;
  hire_date: string;
}

export interface IEmployeeTree extends IEmployeeTreeEl {
  subordinates: IEmployeeTreeEl[];
  colleagues: IEmployeeTreeEl[];
  managers: IEmployeeTreeEl[];
}
