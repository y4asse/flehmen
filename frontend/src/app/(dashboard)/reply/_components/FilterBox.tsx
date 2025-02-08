"use client";
import React from "react";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectGroup,
  SelectItem,
} from "@radix-ui/react-select";
import { ChevronDownIcon } from "lucide-react";
import { useSearchParams, usePathname, useRouter } from "next/navigation";
import { useDebouncedCallback } from "use-debounce";

export const FilterBox = () => {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();
  const [filter, setFilter] = React.useState("最近のからみ");

  const handleFilter = useDebouncedCallback((term: string) => {
    if (term === "recent") {
      setFilter("最近のからみ");
    } else {
      setFilter("ライバル");
    }
    const params = new URLSearchParams(searchParams);
    if (term) {
      params.set("filter", term);
    } else {
      params.delete("filter");
    }
    replace(`${pathname}?${params.toString()}`);
  }, 300);

  return (
    <Select value={filter} onValueChange={handleFilter}>
      <SelectTrigger className="w-[180px] border-2 border-primary rounded-xl text-white p-2 bg-black">
        <SelectValue placeholder="最近のからみ" className="text-white">
          {filter}
        </SelectValue>
        <ChevronDownIcon className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400" />
      </SelectTrigger>
      <SelectContent className="bg-black text-white w-[180px]">
        <SelectGroup>
          <SelectItem
            className="text-white hover:bg-primary mb-2"
            value="recent"
            defaultChecked
          >
            最近のからみ
          </SelectItem>
          <SelectItem className="text-white hover:bg-primary" value="rival">
            ライバル
          </SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};
