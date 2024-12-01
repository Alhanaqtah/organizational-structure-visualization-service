import { Flex, List, Pagination, Typography } from 'antd';
import { MainLayout } from '../main-layout';
import { EmployeeCard } from '../../features/employee-card';
import { IEmployee } from '../../shared/types/employee';
import { $employeesStore, getEmployeesFx } from '../../features/employee-card/model';
import { useEffect } from 'react';
import { useUnit } from 'effector-react';

const CompanyEmployeesList: React.FC = () => {
  const employeeList = useUnit($employeesStore);

  useEffect(() => {
    getEmployeesFx();
  }, []);

  return (
    <MainLayout>
      <List
        style={{ height: '100%' }}
        grid={{ gutter: 16, xs: 1, sm: 2, md: 2, lg: 3, xl: 4, xxl: 5 }}
        dataSource={employeeList ?? []}
        pagination={{
          onChange: (page) => {
            console.log(page);
          },
          pageSize: 10,
        }}
        renderItem={(item) => (
          <List.Item>
            <EmployeeCard employee={item} />
          </List.Item>
        )}
      />
    </MainLayout>
  );
};

export default CompanyEmployeesList;
