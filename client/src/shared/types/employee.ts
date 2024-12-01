export interface IEmployee {
  firstname: string;
  lastname: string;
  role: string;
  city: string;
  id: string;
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
