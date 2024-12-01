import { Background, Controls, ReactFlow, Node, Edge } from '@xyflow/react';
import '@xyflow/react/dist/style.css';
import { useEffect, useState } from 'react';
import { $employeeTreeStore, getEmployeeTreeFx, transformEmployeeTreeToGraph } from './model';
import { useParams } from 'react-router';
import { useUnit } from 'effector-react';

export const CompanyStructureGraph: React.FC = () => {
  const employeeTree = useUnit($employeeTreeStore);
  const [nodes, setNodes] = useState<Node[]>([]);
  const [edges, setEdges] = useState<Edge[]>([]);

  const { id } = useParams();

  useEffect(() => {
    if (id) {
      getEmployeeTreeFx(id);
    }
  }, []);

  useEffect(() => {
    if (employeeTree) {
      const { nodes, edges } = transformEmployeeTreeToGraph(employeeTree);
      setNodes(nodes);
      setEdges(edges);
    }
  }, [employeeTree]);

  return (
    <div style={{ height: '100vh', width: '100%' }}>
      <ReactFlow nodes={nodes} edges={edges}>
        <Background />
        <Controls />
      </ReactFlow>
    </div>
  );
};
