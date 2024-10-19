import { Method } from '@/types/method.type';

export function getPaymentMethod(url: string): Method | undefined {
  const parts = url.split('/');
  const result = parts.length > 2 ? parts[parts.length - 2] : undefined;

  if (result === 'split' || result === 'replacement') {
    return { method: result };
  } else {
    return undefined;
  }
}

export function getPaymentId(url: string) {
  const parts = url.split('/');
  return parts.length > 1 ? parts[parts.length - 1] : undefined;
}
