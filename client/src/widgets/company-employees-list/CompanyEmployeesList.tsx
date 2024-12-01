import { Flex, List, Pagination, Typography } from 'antd';
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
    {
      id: '5',
      firstname: 'Charlie',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Houston',
    },
    {
      id: '6',
      firstname: 'David',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Philadelphia',
    },
    {
      id: '7',
      firstname: 'Eve',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Phoenix',
    },
    {
      id: '8',
      firstname: 'Frank',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'San Antonio',
    },
    {
      id: '9',
      firstname: 'Grace',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'San Diego',
    },
    {
      id: '10',
      firstname: 'Hank',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Dallas',
    },
    {
      id: '11',
      firstname: 'Ivy',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'San Jose',
    },
    {
      id: '12',
      firstname: 'Jack',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Austin',
    },
    {
      id: '13',
      firstname: 'Kelly',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Jacksonville',
    },
    {
      id: '14',
      firstname: 'Larry',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Indianapolis',
    },
    {
      id: '15',
      firstname: 'Marry',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'San Francisco',
    },
    {
      id: '16',
      firstname: 'Nancy',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Columbus',
    },
    {
      id: '17',
      firstname: 'Oscar',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Fort Worth',
    },
    {
      id: '18',
      firstname: 'Patty',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Charlotte',
    },
    {
      id: '19',
      firstname: 'Quincy',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'Detroit',
    },
    {
      id: '20',
      firstname: 'Roger',
      lastname: 'Doe',
      role: 'Software Engineer',
      city: 'El Paso',
    },
  ];

  return (
    <MainLayout>
      <List
        style={{ height: '100%' }}
        grid={{ gutter: 16, xs: 1, sm: 2, md: 2, lg: 3, xl: 4, xxl: 5 }}
        dataSource={employeesList}
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
