import React from "react";
import Link from "next/link";
import Image from "next/image";

// homeItems 配列を定義
const homeItems = [
  { name: "プロフィール", href: "/profile" },
  { name: "FF", href: "/ff" },
  { name: "リプ", href: "/reply" },
  { name: "ひみつ", href: "/secret" },
];

// BackGroundコンポーネント
export const BackGround = () => {
  return (
    <div
      className=" bg-black  before:content-[''] before:absolute before:w-full before:h-full before:bg-[url('/images/cat_bg.svg')] before:bg-center before:bg-no-repeat before:bg-[length:40%] before:opacity-30 before:-z-10"
      style={{
        display: "flex",
        flexDirection: "column",
        width: "100%",
        height: "100vh",
        paddingTop: "4%",
        color: "white",
        position: "relative",
        zIndex: "0",
      }}
    >
      <div className="bgItemContainer" style={styles.bgItemContainer}>
        <Link href="/">
          <Image
            className="hover:animate-noise"
            width={100}
            height={100}
            src="/images/logo.svg"
            style={styles.logo}
            alt="Logo"
          />
        </Link>
        <div className="items" style={styles.items}>
          {/* homeItems 配列をマッピング */}
          {homeItems.map((item) => (
            <HomeItem key={item.name} name={item.name} href={item.href} />
          ))}
        </div>
      </div>
    </div>
  );
};

// HomeItemコンポーネント
type Props = {
  name: string;
  href: string;
};

const HomeItem = ({ name, href }: Props) => {
  return (
    <Link href={href} style={styles.item}>
      <Image
        src="/images/file.svg"
        alt="Icon"
        style={styles.icon}
        width={100}
        height={100}
      />
      <p style={styles.itemName}>{name}</p>
    </Link>
  );
};

export default BackGround;

// CSSスタイル
const styles: { [key: string]: React.CSSProperties } = {
  bg: {
    display: "flex",
    position: "absolute",
    width: "100%",
    height: "80%",
    alignItems: "center",
    justifyContent: "center",
    zIndex: "-1",
  },
  bgCat: {
    position: "absolute",
    display: "flex",
    margin: "4rem 6rem",
    width: "50%",
    height: "auto",
    objectFit: "cover",
    zIndex: "-1",
    opacity: "0.3",
  },
  bgItemContainer: {
    display: "flex",
    flexDirection: "column",
    alignItems: "flex-start",
    justifyContent: "center",
    width: "25rem",
    height: "100%",
  },
  logo: {
    width: "75%",
    height: "auto",
    marginBottom: "2rem",
    margin: "0 5%",
  },
  items: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "flex-start",
    width: "auto",
    height: "100%",
    marginTop: "2rem",
    marginLeft: "10%",
    gap: "4vh",
  },
  item: {
    width: "auto",
    height: "15%",
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    textDecoration: "none",
    color: "white",
    margin: "5%s",
  },
  icon: {
    margin: "2%",
    width: "auto",
    height: "80%",
    marginBottom: "0.5rem",
  },
  itemName: {
    fontSize: "80%",
    textAlign: "center",
    margin: "3%",
    width: "max-content",
  },
};
