import React from "react";
import { content } from "../page";

type Props = {
  repContents?: content[];
};

const Contents = ({ repContents = [] }: Props) => {
  return (
    <div>
      {repContents.map((contents) => (
        <div className="my-4 max-w-lg" key={contents.id}>
          <a className="text-white" href={`https://twitter.com/${contents.id}`}>□@ {contents.id}</a>
          <p className="text-white">{contents.content}</p>
        </div>
      ))}
    </div>
  );
};

export default Contents;
