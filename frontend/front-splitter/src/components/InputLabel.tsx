import React from 'react';

export function InputLabel({ label }: { label: string }) {
  return (
    <label className='block min-w-full pb-2 text-sm font-medium text-gray-300'>
      {label}
    </label>
  );
}
