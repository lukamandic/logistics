import React from "react";
import { Button } from "../atoms/Button";
import { Input } from "../atoms/Input";
import { PackSizeInput } from "../molecules/PackSizeInput";
import { DistributionResultsTable } from "../molecules/DistributionResultsTable";
import { usePackSizes } from "../../hooks/usePackSizes";
import { useDistribution } from "../../hooks/useDistribution";

export const Calculator: React.FC = () => {
  const {
    packSizes,
    setPackSizes,
    isLoading: isPackSizesLoading,
    error: packSizesError,
    savePackSizes,
  } = usePackSizes();

  const {
    amount,
    setAmount,
    distribution,
    isLoading: isDistributionLoading,
    error: distributionError,
    calculateDistribution,
    resetDistribution,
  } = useDistribution();

  const handleSubmitPackSizes = async () => {
    await savePackSizes();
    resetDistribution();
  };

  const isLoading = isPackSizesLoading || isDistributionLoading;
  const error = packSizesError || distributionError;

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-4xl mx-auto">
        <div className="text-center mb-12">
          <h1 className="text-4xl font-bold text-gray-900 mb-4">
            Order Packs Calculator
          </h1>
          <p className="text-lg text-gray-600">
            Calculate the optimal distribution of pack sizes for your order
          </p>
        </div>

        <div className="space-y-8">
          {/* Pack Sizes Section */}
          <section className="bg-white rounded-2xl shadow-lg p-8 transition-all hover:shadow-xl">
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-2xl font-semibold text-gray-900">
                Pack Sizes
              </h2>
              <Button
                onClick={handleSubmitPackSizes}
                disabled={isLoading || packSizes.length === 0}
                className="ml-4"
              >
                {isLoading ? "Saving..." : "Save Pack Sizes"}
              </Button>
            </div>
            <PackSizeInput
              packSizes={packSizes}
              onPackSizesChange={setPackSizes}
            />
          </section>

          {/* Calculate Section */}
          <section className="bg-white rounded-2xl shadow-lg p-8 transition-all hover:shadow-xl">
            <h2 className="text-2xl font-semibold text-gray-900 mb-6">
              Calculate Distribution
            </h2>

            <div className="flex flex-col sm:flex-row gap-4 items-end mb-6">
              <div className="flex-1">
                <Input
                  type="number"
                  value={amount}
                  onChange={setAmount}
                  label="Order Amount"
                  placeholder="Enter the number of items needed"
                  min={0}
                  required
                  className="w-full"
                />
              </div>
              <Button
                onClick={calculateDistribution}
                disabled={isLoading || !amount || packSizes.length === 0}
                className="whitespace-nowrap"
              >
                {isLoading ? "Calculating..." : "Calculate"}
              </Button>
            </div>

            <DistributionResultsTable distribution={distribution} />
          </section>

          {/* Error Message */}
          {error && (
            <div className="bg-red-50 border-l-4 border-red-500 p-4 rounded-lg">
              <div className="flex">
                <div className="flex-shrink-0">
                  <svg
                    className="h-5 w-5 text-red-500"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fillRule="evenodd"
                      d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                      clipRule="evenodd"
                    />
                  </svg>
                </div>
                <div className="ml-3">
                  <p className="text-sm text-red-700">{error}</p>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
};
