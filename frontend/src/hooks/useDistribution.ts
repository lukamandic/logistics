import { useState } from "react";
import { API_BASE_URL } from "../config/env";

interface UseDistributionReturn {
  amount: string;
  setAmount: React.Dispatch<React.SetStateAction<string>>;
  distribution: Record<string, number>;
  isLoading: boolean;
  error: string | null;
  calculateDistribution: () => Promise<void>;
  resetDistribution: () => void;
}

export const useDistribution = (): UseDistributionReturn => {
  const [amount, setAmount] = useState<string>("");
  const [distribution, setDistribution] = useState<Record<string, number>>({});
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const calculateDistribution = async () => {
    try {
      setIsLoading(true);
      setError(null);

      const response = await fetch(
        `${API_BASE_URL}/calculate?amount=${amount}`,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      if (!response.ok) {
        throw new Error("Failed to calculate distribution");
      }

      const data = await response.json();
      setDistribution(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred");
    } finally {
      setIsLoading(false);
    }
  };

  const resetDistribution = () => {
    setDistribution({});
  };

  return {
    amount,
    setAmount,
    distribution,
    isLoading,
    error,
    calculateDistribution,
    resetDistribution,
  };
};
