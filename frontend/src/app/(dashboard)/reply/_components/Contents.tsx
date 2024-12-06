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
          <p className="text-white">□@ {contents.id}</p>
          <p className="text-white">{contents.content}</p>
        </div>
      ))}
    </div>
  );
};

export default Contents;
