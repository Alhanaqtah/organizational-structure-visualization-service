import { Button, Form, Layout } from 'antd';
import { Content, Header } from 'antd/es/layout/layout';
import Sider from 'antd/es/layout/Sider';
import { PropsWithChildren, useState } from 'react';
import './MainLayout.css';
import { MainHeader } from '../header';
import { Filter } from '../../features';

interface MainLayoutProps extends PropsWithChildren {}

export const MainLayout: React.FC<MainLayoutProps> = ({ children }) => {
  const [form] = Form.useForm();

  return (
    <Layout className="mainLayout">
      <Form form={form}>
        <Header className="headerLayout">
          <MainHeader />
        </Header>
        <Layout>
          <Sider className="siderLayout" width="250">
            <Filter />
          </Sider>
          <Content className="contentLayout">{children}</Content>
        </Layout>
      </Form>
    </Layout>
  );
};
