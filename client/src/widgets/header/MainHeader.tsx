import { SearchOutlined } from '@ant-design/icons';
import { Button, Flex, Input, Tooltip } from 'antd';
import { Logo } from '../../shared/ui';

interface HeaderProps {}

export const MainHeader: React.FC<HeaderProps> = () => {
  return (
    <Flex style={{ padding: '8px 16px', height: '100%' }} gap={16} align="center" justify="center">
      <div style={{ width: '218px', height: '82px', flexShrink: 0 }}>
        <Logo />
      </div>
      <Input placeholder="Фамилия" allowClear style={{ height: '46px' }} />
      <Input placeholder="Имя" allowClear style={{ height: '46px' }} />
      <Input placeholder="Отчество" allowClear style={{ height: '46px' }} />
      <Tooltip title="Поиск">
        <Button
          shape="circle"
          type="primary"
          icon={<SearchOutlined />}
          style={{
            flexShrink: 0,
            height: '46px',
            width: '46px',
            borderRadius: '10px',
          }}
        />
      </Tooltip>
    </Flex>
  );
};
