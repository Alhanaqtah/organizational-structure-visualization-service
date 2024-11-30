import { Button, Layout } from 'antd';
import { Content, Header } from 'antd/es/layout/layout';
import Sider from 'antd/es/layout/Sider';
import { PropsWithChildren, useState } from 'react';
import './MainLayout.css';
import { MainHeader } from '../header';
import Icon, { MenuFoldOutlined, MenuUnfoldOutlined } from '@ant-design/icons';
import { Logo } from '../../shared/ui';

interface MainLayoutProps extends PropsWithChildren {}

export const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  const [collapsed, setCollapsed] = useState(false);
  return (
    <Layout className="mainLayout">
      <Header className="headerLayout">
        <MainHeader />
      </Header>
      <Layout>
        <Sider className="siderLayout" width="250"></Sider>
        <Content className="contentLayout">{children}</Content>
      </Layout>
    </Layout>
  );
};
