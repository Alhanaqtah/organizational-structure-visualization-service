import './index.css';
import { CompanyEmployeesListPage } from '../pages/company-employees-list/CompanyEmployeesListPage';
import { ThemeProvider } from './providers/ThemeProvider';
// import CompanyStructurePage from './pages/company-structure/CompanyStructurePage';

function App() {
  return (
    <div id="app" style={{ width: '100vw', height: '100wh' }}>
      <ThemeProvider>
        <CompanyEmployeesListPage />
      </ThemeProvider>
    </div>
  );
}

export default App;
