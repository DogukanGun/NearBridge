import React, { useState } from 'react';

export const Checkbox = (props: CheckboxProps) => {
  const [isChecked, setIsChecked] = useState<boolean>(false);

  const handleCheckboxChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { checked } = event.target;
    setIsChecked(checked);
    props.onChecked(checked);
  };

  return (
    <div className="flex items-center my-10 mx-auto">
      <input
        checked={isChecked}
        onChange={handleCheckboxChange}
        id="default-checkbox"
        type="checkbox"
        className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
      />
      <label
        htmlFor="default-checkbox"
        className="ml-2 text-sm font-medium text-black"
      >
        {props.label}
      </label>
    </div>
  );
};

interface CheckboxProps {
  label: string;
  onChecked: (checked: boolean) => void;
}
