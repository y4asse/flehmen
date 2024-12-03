import React, { useState, useEffect } from 'react';

interface TypewriterProps {
  text: string;
  speed?: number;
  loop?: boolean;
}

const Typewriter: React.FC<TypewriterProps> = ({ text, speed = 100, loop = false }) => {
  const [displayedText, setDisplayedText] = useState('');
  const [index, setIndex] = useState(0);

  useEffect(() => {
    if (index < text.length) {
      const timer = setTimeout(() => {
        setDisplayedText((prev) => prev + text.charAt(index));
        setIndex((prev) => prev + 1);
      }, speed);
      return () => clearTimeout(timer);
    } else if (loop) {
      setTimeout(() => {
        setDisplayedText('');
        setIndex(0);
      }, speed);
    }
  }, [index, text, speed, loop]);

  return (
    <div style={{ position: 'relative', whiteSpace: 'pre' }}>
      {/* 完成形のテキストを透明で表示 */}
      <div style={{ visibility: 'hidden' }}>{text}</div>
      {/* 実際のタイプライター効果 */}
      <div style={{ position: 'absolute', top: 0, left: 0 }}>
        {displayedText}
      </div>
    </div>
  );
};

export default Typewriter;
