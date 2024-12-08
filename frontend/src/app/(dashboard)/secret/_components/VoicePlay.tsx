"use client";
import React, { useState, useRef } from "react";
import { Flex } from "@/components/ui/flex";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import Image from "next/image";

const AudioPlayer = () => {
  const [isPlaying, setIsPlaying] = useState(false);
  const [progress, setProgress] = useState(0);
  const audioRef = useRef<HTMLAudioElement>(null);

  const handlePlayPause = () => {
    if (!audioRef.current) return;
    if (isPlaying) {
      audioRef.current.pause();
    } else {
      audioRef.current.play();
    }
    setIsPlaying(!isPlaying);
  };

  const handleSeek = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (!audioRef.current) return;
    const time = Number(e.target.value);
    const seekTime = (time / 100) * audioRef.current.duration;
    audioRef.current.currentTime = seekTime;
    setProgress(time);
  };

  const updateProgress = () => {
    if (audioRef.current) {
      const progressPercent =
        (audioRef.current.currentTime / audioRef.current.duration) * 100;
      setProgress(progressPercent);
    }
  };

  const handleEnd = () => {
    setIsPlaying(false);
    setProgress(0);
  };

  return (
    <Flex className="relative flex flex-col items-center space-y-4 p-4 w-full h-[400px]">
      <div className="relative inline-block">
        <Image
          src="/images/mic.svg"
          alt="Mic"
          width={100}
          height={100}
          className="transition-transform duration-300 hover:scale-110"
        />
      </div>

      <Button
        onClick={handlePlayPause}
        className="text-white p-4 rounded-full bg-gray-800"
        style={{
          width: "20%",
          height: "20%",
          fontSize: "1.5rem",
          imageRendering: "pixelated",
        }}
      >
        {isPlaying ? "❚❚" : "▷"}
      </Button>

      <Input
        type="range"
        min="0"
        max="100"
        value={progress}
        onChange={handleSeek}
        className="w-3/4 appearance-none bg-gray-300 rounded-lg outline-none"
        style={{
          width: "70%",
          height: "10px",
          background: `linear-gradient(to right, #E4007F ${progress}%, #ccc ${progress}%)`,
        }}
      />

      {/* ダウンロードボタン */}
      <Button
        className="absolute bottom-4 right-4   p-4 rounded-full "
        onClick={() => {
          alert("ダウンロードボタンがクリックされました");
        }}
      >
        ↓ダウンロード
      </Button>

      <style jsx>{`
        input[type="range"]::-webkit-slider-thumb {
          -webkit-appearance: none;
          appearance: none;
          width: 32px;
          height: 32px;
          background: url("/images/hart_item.png") no-repeat center;
          background-size: contain;
          cursor: pointer;
          image-rendering: pixelated;
          transform: translateY(-50%);
        }
        input[type="range"]::-moz-range-thumb {
          width: 32px;
          height: 32px;
          background: url("/images/hart_item.png") no-repeat center;
          background-size: contain;
          cursor: pointer;
          image-rendering: pixelated;
        }
      `}</style>

      <audio
        ref={audioRef}
        src="./voice/yase_voice.mp3"
        onTimeUpdate={updateProgress}
        onEnded={handleEnd}
      />
    </Flex>
  );
};

export default AudioPlayer;
