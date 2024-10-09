import { NextResponse } from 'next/server';
import { z } from 'zod';

import { axiosInstance } from '@/utils/axiosInstance';

const paymentRequestSchema = z.object({
  groupId: z.string(),
});

export async function GET(request: Request) {
  try {
    console.log('tetetetete');
    const url = new URL(request.url);
    const groupId = url.searchParams.get('groupId');

    if (!groupId) {
      return NextResponse.json(
        { message: 'Invalid request body' },
        { status: 400 },
      );
    }

    const parsedBody = paymentRequestSchema.safeParse({ groupId });

    if (!parsedBody.success) {
      return NextResponse.json(
        { message: 'Invalid request body' },
        { status: 400 },
      );
    }

    const response = await axiosInstance.get(
      'payments/' + parsedBody.data.groupId,
    );

    return NextResponse.json(response.data, { status: 200 });
  } catch (error) {
    console.error('Failed to fetch:', error);
    return NextResponse.json(
      { message: 'Failed to fetch group-data' },
      { status: 500 },
    );
  }
}
