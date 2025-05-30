import React from "react";
import { Input } from "../atoms/Input";
import { Button } from "../atoms/Button";

interface PackSizeInputProps {
  packSizes: number[];
  onPackSizesChange: (sizes: number[]) => void;
}

export const PackSizeInput: React.FC<PackSizeInputProps> = ({
  packSizes,
  onPackSizesChange,
}) => {
  const handleAddPackSize = () => {
    onPackSizesChange([...packSizes, 0]);
  };

  const handleRemovePackSize = (index: number) => {
    const newPackSizes = packSizes.filter((_, i) => i !== index);
    onPackSizesChange(newPackSizes);
  };

  const handlePackSizeChange = (index: number, value: string) => {
    const newValue = parseFloat(value) || 0;
    const newPackSizes = [...packSizes];
    newPackSizes[index] = newValue;
    onPackSizesChange(newPackSizes);
  };

  return (
    <div className="space-y-4">
      <div className="space-y-3 flex flex-col items-center">
        {packSizes.map((size, index) => (
          <div
            key={index}
            className="flex justify-center w-[50%] items-end gap-2 bg-gray-50 p-4 rounded-lg"
          >
            <Input
              type="number"
              value={size}
              onChange={(value) => handlePackSizeChange(index, value)}
              label={`Pack Size ${index + 1}`}
              min={0}
              required
              className="flex-1"
            />
            <Button
              variant="secondary"
              onClick={() => handleRemovePackSize(index)}
              className="mb-1 px-3 py-2"
            >
              <svg className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path
                  fillRule="evenodd"
                  d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                  clipRule="evenodd"
                />
              </svg>
            </Button>
          </div>
        ))}
      </div>
      <Button
        onClick={handleAddPackSize}
        variant="secondary"
        className="w-full"
      >
        <svg
          className="h-5 w-5 mr-2 inline-block"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fillRule="evenodd"
            d="M10 3a1 1 0 011 1v5h5a1 1 0 110 2h-5v5a1 1 0 11-2 0v-5H4a1 1 0 110-2h5V4a1 1 0 011-1z"
            clipRule="evenodd"
          />
        </svg>
        Add Pack Size
      </Button>
    </div>
  );
};
