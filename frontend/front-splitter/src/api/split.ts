import { SplitResponse } from '@/types/paymentDetailes/splitResponse.type';
import { axiosInstance } from '@/utils/axiosInstance';
import axios from 'axios';

export async function getSplit(splitId: string) {
  try {
    const res = await axiosInstance.get<SplitResponse>(`splits/${splitId}`, {
      headers: {
        'Cache-Control': 'no-store',
      },
    });

    return res.data;
  } catch (e) {
    if (axios.isAxiosError(e)) {
      console.error(e.message);
    }
    return null;
  }
}
