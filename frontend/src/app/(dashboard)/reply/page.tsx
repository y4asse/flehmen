import { Suspense } from "react";
import ReplyPage from "./_components/ReplyPage";

const Page = () => {
  return (
    <Suspense fallback={<div>loading...</div>}>
      <ReplyPage />
    </Suspense>
  );
};

export default Page;
