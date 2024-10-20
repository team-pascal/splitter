import { HistoryResponse } from '@/types/history.type';
import { axiosInstance } from '@/utils/axiosInstance';
import axios from 'axios';

export async function getPayment(groupId: string): Promise<HistoryResponse> {
  try {
    const response = await axiosInstance.get<HistoryResponse>(
      `/payments/${groupId}`,
      {
        headers: {
          'Cache-Control': 'no-store',
        },
      },
    );
    return response.data;
  } catch (e) {
    if (axios.isAxiosError(e)) {
      console.error(e.message);
    }
    return [];
  }
}
