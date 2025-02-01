import React from "react";
import Image from "next/image";
import { Icon } from "../page";

type Props = {
  Icon: Icon[];
};

export const HomeIcon = (props: Props) => {
  const { Icon } = props;
  console.log(Icon.map((icon) => icon.image));
  return (
    <div
      style={{
        position: "absolute",
        top: "0",
        left: "0",
        width: "100vw",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        marginTop: "70px",
      }}
    >
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "1fr 1fr",
          gap: "50px",
        }}
      >
        {Icon.map((icon) => (
          <a key={icon.name} href={icon.href}>
            <Image src={icon.image} alt={icon.name} width={140} height={100} />
          </a>
        ))}
      </div>
    </div>
  );
};
