import { NextResponse } from 'next/server';

import { axiosInstance } from '@/utils/axiosInstance';

export async function GET(request: Request) {
  try {
    const url = new URL(request.url);
    const groupId = url.searchParams.get('groupId');

    if (!groupId) {
      return NextResponse.json(
        { message: 'Invalid request body' },
        { status: 400 },
      );
    }

    const response = await axiosInstance.get('payments/' + groupId);

    return NextResponse.json(response.data, { status: 200 });
  } catch {
    return NextResponse.json(
      { message: 'Failed to fetch group-data' },
      { status: 500 },
    );
  }
}
