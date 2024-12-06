import React from "react";
import { reply } from "../page";

type rep = {
  replyInfo?: reply[];
};

const ReplyList = ({ replyInfo = [] }: rep = {}) => {
  return (
    <div>
      {replyInfo.map((reply, index) => (
        <div className="my-4" key={`${reply.id}-${index}`}>
          <a className="text-white" href={`https://twitter.com/${reply.id}`}>□@ {reply.id} </a>
          <p className="text-white">過去1週間でリプした回数: {reply.count}</p>
          <p className="text-white">親密度: {reply.intimacy}</p>
        </div>
      ))}
    </div>
  );
};

export default ReplyList;
