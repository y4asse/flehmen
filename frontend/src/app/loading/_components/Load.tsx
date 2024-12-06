'use client'; // クライアントコンポーネントとしてマーク

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';

const Load = () => {
  const router = useRouter(); // Next.jsのルーターを使用

  useEffect(() => {
    // 4秒後にホーム画面へ遷移
    const timer = setTimeout(() => {
      router.push('/');
    }, 2000);

    // クリーンアップ関数
    return () => clearTimeout(timer);
  }, [router]);

  return (
    <div className="relative w-screen h-screen bg-black flex items-center justify-center overflow-hidden">
      {/* ロゴ表示 */}
      <div className="z-10 animate-pulse ">
        <img
          src="/images/logo.gif"
          alt="Loading Logo"
          className="w-90 h-auto"
        />
      </div>

      {/* ノイズエフェクト */}
      <div className="absolute inset-0 z-0 opacity-20">
        <div className="noise-bg w-full h-full"></div>
      </div>

      {/* カスタムスタイル */}
      <style jsx>{`
        .noise-bg {
          background-image: url('/images/noise.png'); // ノイズ画像を適用
          background-size: cover;
          animation: noiseAnimation 0.2s infinite;
        }

        @keyframes noiseAnimation {
          0% { transform: translate(0, 0); }
          25% { transform: translate(-5%, 5%); }
          50% { transform: translate(5%, -5%); }
          75% { transform: translate(-5%, -5%); }
          100% { transform: translate(5%, 5%); }
        }
      `}</style>
    </div>
  );
};

export default Load;
