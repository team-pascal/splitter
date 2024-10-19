export function formatDate(dateStr: string) {
  const date = new Date(dateStr);

  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDay()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');

  return `${year}年${month}月${day} ${hours}:${minutes}`;
}
