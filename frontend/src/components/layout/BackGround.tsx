import React from "react";
import Link from "next/link";
import Image from "next/image";

// homeItems 配列を定義
const homeItems = [
  { name: "生活習慣", href: "/life" },
  { name: "おともだち", href: "/friendship" },
  { name: "ff", href: "/ff" },
  { name: "ひみつ", href: "/secret" },
];

// BackGroundコンポーネント
export const BackGround = () => {
  return (
    <div style={styles.backGround}>
      <Link href={"/"} className="bg" style={styles.bg}>
        <Image
          className="bg_cat"
          width={100}
          height={100}
          src="/images/cat_bg.svg"
          style={styles.bgCat}
          alt="Background Cat"
        />
      </Link>

      <div className="logoContainer" style={styles.logoContainer}>
        <Image
          className="logo"
          width={100}
          height={100}
          src="/images/logo.svg"
          style={styles.logo}
          alt="Logo"
        />
      </div>

      <div style={styles.itemContainer}>
        {/* homeItems 配列をマッピング */}
        {homeItems.map((item) => (
          <HomeItem key={item.name} name={item.name} href={item.href} />
        ))}
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
        src="/images/file_icon.png"
        alt="Icon"
        style={styles.icon}
        width={100}
        height={100}
      />
      <div style={styles.itemName}>{name}</div>
    </Link>
  );
};

export default BackGround;

// CSSスタイル
const styles: { [key: string]: React.CSSProperties } = {
  backGround: {
    backgroundColor: "black",
    display: "flex",
    flexDirection: "column",
    width: "100%",
    minHeight: "100vh",
    paddingTop: "4rem",
    fontFamily: '"PixelMplus10"',
    color: "white",
    position: "relative", // 背景画像のレイヤー制御のために追加
    zIndex: "0", // 背景画像のレイヤー制御のために追加
  },
  bg: {
    display: "flex",
    position: "absolute",
    width: "100%",
    height: "80%",
    alignItems: "center",
    justifyContent: "center",
    marginTop: "2rem",
  },
  bgCat: {
    position: "absolute",
    display: "flex",
    margin: "4rem 6rem",
    width: "50%",
    height: "auto",
    objectFit: "cover", // 背景画像がコンテナに合わせてカバーするように調整
    zIndex: "-1", // 背景を後ろに配置
    opacity: "0.3", // 背景画像を透過
  },
  logoContainer: {
    margin: "0rem 2rem",
    display: "flex",
  },
  logo: {
    width: "auto",
    height: "7rem",
    marginBottom: "2rem",
  },
  itemContainer: {
    //logoContainerの下に配置
    position: "absolute",
    display: "flex",
    left: "3rem",
    top: "12rem",
    flexDirection: "column",
    alignItems: "center", // 左寄せにする
    padding: "0 1rem",
    gap: "2rem",
  },
  item: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    textDecoration: "none",
    color: "white",
  },
  icon: {
    width: "5rem",
    height: "auto",
    marginBottom: "0.5rem",
  },
  itemName: {
    fontSize: "20px",
    textAlign: "center", // テキストを中央揃え
  },
};
