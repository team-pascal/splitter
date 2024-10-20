import { useState, useEffect } from 'react';

import { SplitResponse } from '@/types/paymentDetailes/splitResponse.type';
import { getSplit } from '@/api/split';
import axios from 'axios';

export function usePaymentDetail(paymentId: string | undefined) {
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<SplitResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      try {
        const id = paymentId ? paymentId : '';
        const res = await getSplit(id);
        setData(res);
      } catch (e) {
        if (axios.isAxiosError(e)) {
          console.error(e);
          setError(e.message);
        }
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, []);
  return [loading, data, error];
}
