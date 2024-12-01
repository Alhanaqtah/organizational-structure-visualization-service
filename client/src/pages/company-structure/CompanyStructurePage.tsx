import { Outlet } from 'react-router';
import { CompanyStructureGraph } from '../../widgets/company-structure-graph';

const CompanyStructurePage: React.FC = () => {
  return (
    <>
      <CompanyStructureGraph />
      <Outlet />
    </>
  );
};

export default CompanyStructurePage;
