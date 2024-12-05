import { Flex } from "@/components/ui/flex";
import { Search } from "@/components/ui/search";
import React, { Suspense } from "react";
import UniversityList from "./UniversityList";

type Props = {
  query: string;
};

export const UniversitySearch = (props: Props) => {
  const { query } = props;

  return (
    <Flex
      className="h-full w-full p-4 gap-4"
      align={"start"}
      justify={"start"}
      direction={"column"}
    >
      <Search placeholder="略称もいけるよ" />
      <Suspense fallback={<div>loading...</div>}>
        <UniversityList query={query} />
      </Suspense>
    </Flex>
  );
};
