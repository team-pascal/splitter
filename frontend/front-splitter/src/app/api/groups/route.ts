import { NextResponse } from 'next/server';
import { z } from 'zod';

import { axiosInstance } from '@/utils/axiosInstance';

const groupRequestSchema = z.object({
  name: z.string(),
  users: z.array(z.string()),
});

export async function POST(request: Request) {
  try {
    const body: unknown = await request.json();

    const parsedBody = groupRequestSchema.safeParse(body);

    if (!parsedBody.success) {
      return NextResponse.json(
        { message: 'Invalid request body' },
        { status: 400 },
      );
    }

    const response = await axiosInstance.post('/groups', parsedBody.data);
    return NextResponse.json(response.data, { status: 200 });
  } catch (error) {
    console.error('Failed to fetch:', error);
    return NextResponse.json(
      { message: 'Failed to create grouo' },
      { status: 500 },
    );
  }
}
