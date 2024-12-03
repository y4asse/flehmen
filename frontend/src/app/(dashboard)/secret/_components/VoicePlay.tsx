"use client";
import React, { useState, useRef, useEffect } from 'react';
import { Play, Pause } from 'lucide-react';

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
    <div className="flex items-center space-x-4 p-4 bg-gray-100 rounded-lg">
      <button 
        onClick={handlePlayPause} 
        className="bg-blue-500 text-white p-2 rounded-full hover:bg-blue-600"
      >
        {isPlaying ? <Pause size={24} /> : <Play size={24} />}
      </button>
      
      <input 
        type="range"
        min="0"
        max="100"
        value={progress}
        onChange={handleSeek}
        className="flex-grow"
      />
      
      <audio 
        ref={audioRef} 
        src="./voice/ayapo_voice.mp3"
        onTimeUpdate={updateProgress}
        onEnded={handleEnd}
      />
    </div>
  );
};

export default AudioPlayer;