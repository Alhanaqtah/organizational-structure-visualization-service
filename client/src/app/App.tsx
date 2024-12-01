import './index.css';
import { CompanyEmployeesListPage } from '../pages/company-employees-list/CompanyEmployeesListPage';
import { ThemeProvider } from './providers/ThemeProvider';
import { BrowserRouter, Route, Routes } from 'react-router';
import CompanyStructurePage from '../pages/company-structure/CompanyStructurePage';
import ErrorBoundary from 'antd/es/alert/ErrorBoundary';
// import CompanyStructurePage from './pages/company-structure/CompanyStructurePage';

function App() {
  return (
    <div id="app" style={{ width: '100vw', height: '100wh' }}>
      <ErrorBoundary>
        <ThemeProvider>
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<CompanyEmployeesListPage />}>
                <Route path="info/:id" element={<CompanyStructurePage />} />
              </Route>
            </Routes>
          </BrowserRouter>
        </ThemeProvider>
      </ErrorBoundary>
    </div>
  );
}

export default App;
