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
        <Image
          className="bg_cat"
          width={100}
          height={100}
          src="/images/cat_bg.svg"
          style={styles.bgCat}
          alt="Background Cat"
        />
      <div className="bgItemContainer"style={styles.bgItemContainer}>
        <Image
          className="logo"
          width={100}
          height={100}
          src="/images/logo.svg"
          style={styles.logo}
          alt="Logo"
        />
        <div className= "items" style = {styles.items}>
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
    height: "100vh",
    paddingTop: "4%",
    color: "white",
    position: "relative", 
    zIndex: "0", 
  },
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
    width: "100%",
    height: "auto",
    marginBottom: "2rem",
    margin: "0 5%",
  },
  items:{
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "flex-start",
    width: "auto",
    height: "100%",
    marginTop: "2rem",
    marginLeft: "10%",
    gap: "2vh",
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
    margin:"2%" ,
    width: "auto",
    height: "80%",
    marginBottom: "0.5rem",
  },
  itemName: {
    fontSize: "80%",
    textAlign: "center", 
    margin:"3%"
  },
};
