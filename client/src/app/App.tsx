import './index.css';
import { CompanyEmployeesListPage } from '../pages/company-employees-list/CompanyEmployeesListPage';
import { ThemeProvider } from './providers/ThemeProvider';
import { BrowserRouter, Route, Routes } from 'react-router';
import CompanyStructurePage from '../pages/company-structure/CompanyStructurePage';
// import CompanyStructurePage from './pages/company-structure/CompanyStructurePage';

function App() {
  return (
    <div id="app" style={{ width: '100vw', height: '100wh' }}>
      <ThemeProvider>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<CompanyEmployeesListPage />}>
              <Route path="info/:id" element={<CompanyStructurePage />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </ThemeProvider>
    </div>
  );
}

export default App;
