import { ConfigProvider as AntdConfigProvider, theme } from 'antd';
import { PropsWithChildren } from 'react';

interface ThemeProviderProps extends PropsWithChildren {}
export const ThemeProvider: React.FC<ThemeProviderProps> = ({ children }) => {
  return (
    <AntdConfigProvider
      theme={{
        components: {
          Button: {
            colorPrimary: '#ED1D24',
            algorithm: true,
          },
          Pagination: {
            colorPrimary: '#ED1D24',
            colorPrimaryHover: '#ED1D24',
          },
        },
      }}
    >
      {children}
    </AntdConfigProvider>
  );
};
