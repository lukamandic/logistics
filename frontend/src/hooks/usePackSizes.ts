import { useState, useEffect } from "react";
import { API_BASE_URL } from "../config/env";

interface UsePackSizesReturn {
  packSizes: number[];
  setPackSizes: React.Dispatch<React.SetStateAction<number[]>>;
  isLoading: boolean;
  error: string | null;
  savePackSizes: () => Promise<void>;
}

export const usePackSizes = (): UsePackSizesReturn => {
  const [packSizes, setPackSizes] = useState<number[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchParcelSizes = async () => {
      try {
        setIsLoading(true);
        setError(null);

        const response = await fetch(`${API_BASE_URL}/parcel-sizes`);
        if (!response.ok) {
          if (response.status === 404) {
            // No parcel sizes configured yet, this is not an error
            return;
          }
          throw new Error("Failed to fetch pack sizes");
        }

        const data = await response.json();
        if (data && data.length > 0) {
          // Use the first item's parcel sizes
          setPackSizes(data[0].parcel_sizes);
        }
      } catch (err) {
        setError(err instanceof Error ? err.message : "An error occurred");
      } finally {
        setIsLoading(false);
      }
    };

    fetchParcelSizes();
  }, []);

  const savePackSizes = async () => {
    try {
      setIsLoading(true);
      setError(null);

      const response = await fetch(`${API_BASE_URL}/parcel-sizes`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: "1",
          parcel_sizes: packSizes,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to update pack sizes");
      }
    } catch (err) {
      setError(err instanceof Error ? err.message : "An error occurred");
    } finally {
      setIsLoading(false);
    }
  };

  return {
    packSizes,
    setPackSizes,
    isLoading,
    error,
    savePackSizes,
  };
};
