import { Checkbox, Space, Typography, Button } from 'antd';
import React, { useEffect } from 'react';
import { $filtersStore, getFiltersFx } from './model';
import { useUnit } from 'effector-react';
import { $activeFiltersStore, getEmployeesFx, QueryParams, updateFilter } from '../employee-card/model';

interface FilterProps {}

export const Filter: React.FC<FilterProps> = () => {
  const filtersStore = useUnit($filtersStore);
  const activeFiltersStore = useUnit($activeFiltersStore);

  useEffect(() => {
    getFiltersFx();
  }, []);

  const handleCheckboxChange = (filter: keyof QueryParams, value: string) => {
    updateFilter({ filter, value });
  };

  const handleApplyFilters = () => {
    getEmployeesFx();
  };

  return (
    <>
      {filtersStore.map((data, id) => (
        <Space direction="vertical" key={id}>
          <Typography style={{ fontWeight: 'bold', fontSize: '25px' }}>{data.name}</Typography>
          <div style={{ display: 'flex', flexWrap: 'wrap' }}>
            {data.values.map((value) => {
              const filterValues = activeFiltersStore[data.filter as keyof QueryParams];
              const isChecked = Array.isArray(filterValues) && filterValues.includes(value);
              return (
                <Checkbox
                  key={value}
                  checked={isChecked}
                  onChange={() => handleCheckboxChange(data.filter as keyof QueryParams, value)}
                >
                  {value}
                </Checkbox>
              );
            })}
          </div>
        </Space>
      ))}
      <Button type="primary" onClick={handleApplyFilters}>
        Применить фильтры
      </Button>
    </>
  );
};
