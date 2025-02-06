import { Flex } from "@/components/ui/flex";
import { Input } from "@/components/ui/input";
import React from "react";

type Props = {
  handleFileChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
};
export const FileUploader = (props: Props) => {
  const { handleFileChange } = props;

  return (
    <Flex direction={"column"} className="p-4 gap-6">
      <label htmlFor="file" className="text-white">
        LINEのメッセージをアップロードしよう
      </label>
      <Input
        type="file"
        onChange={handleFileChange}
        id="file"
        accept=".txt"
        className="text-white"
      />
    </Flex>
  );
};
