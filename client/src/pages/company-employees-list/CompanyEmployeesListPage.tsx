import { Outlet } from 'react-router';
import CompanyEmployeesList from '../../widgets/company-employees-list/CompanyEmployeesList';

export const CompanyEmployeesListPage: React.FC = () => {
  return (
    <>
      <Outlet />;
      <CompanyEmployeesList />
    </>
  );
};
