import React from "react";
import Image from "next/image";
import { Icon } from "../page";
import Link from "next/link";

type Props = {
  Icon: Icon[];
};

export const HomeIcon = (props: Props) => {
  const { Icon } = props;
  return (
    <div
      style={{
        display: "grid",
        gridTemplateColumns: "1fr 1fr",
        width: "100%",
        height: "100%",
        position: "relative",
      }}
    >
      {Icon.map((icon) => (
        <Link
          key={icon.name}
          href={icon.href}
          className="flex flex-col justify-center items-center gap-2 "
        >
          <Image
            src={icon.image}
            alt={icon.name}
            width={150}
            height={150}
            className="absolute"
          />
          <p className="text-white relative top-16">{icon.name}</p>
        </Link>
      ))}
    </div>
  );
};
