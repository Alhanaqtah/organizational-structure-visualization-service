import { Button, Card, Space, Typography } from 'antd';
import { IEmployee } from '../../shared/types/employee';

interface EmployeeCardProps {
  employee: IEmployee;
}
export const EmployeeCard: React.FC<EmployeeCardProps> = ({ employee }) => {
  return (
    <Card>
      <Space direction="vertical" style={{ width: '100%' }}>
        <Typography>{`${employee.firstname} ${employee.lastname}`}</Typography>
        <Typography>{employee.role}</Typography>
        <Typography>{employee.city}</Typography>
        <Button style={{ width: '100%', height: '40px' }} type="primary">
          Подробнее
        </Button>
      </Space>
    </Card>
  );
};
