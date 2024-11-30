import { Flex, List, Typography } from 'antd';
import { MainLayout } from '../main-layout';
import { EmployeeCard } from '../../features/employee-card';
import { IEmployee } from '../../shared/types/employee';

const CompanyEmployeesList: React.FC = () => {
  const employeesList: IEmployee[] = [
    {
      id: '1',
      firstname: 'John',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'New York',
    },
    {
      id: '2',
      firstname: 'Jane',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'San Francisco',
    },
    {
      id: '3',
      firstname: 'Alice',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Los Angeles',
    },
    {
      id: '4',
      firstname: 'Bob',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Chicago',
    },
  ];

  return (
    <MainLayout>
      <List
        grid={{ gutter: 16, xs: 1, sm: 2, md: 2, lg: 3, xl: 4, xxl: 5 }}
        dataSource={employeesList}
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
