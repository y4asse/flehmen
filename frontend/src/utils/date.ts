export function formatDateString(dateString: string): string {
  const date = new Date(dateString);

  const year = date.getFullYear();
  const month = date.getMonth() + 1; // 月は0から始まるため+1
  const day = date.getDate();
  const hours = date.getHours();
  const minutes = date.getMinutes();
  const seconds = date.getSeconds();

  // ゼロ埋め関数
  const zeroPad = (num: number): string => num.toString().padStart(2, "0");

  return `${year}年${month}月${day}日 ${hours}時${zeroPad(minutes)}分${zeroPad(
    seconds
  )}秒`;
}
