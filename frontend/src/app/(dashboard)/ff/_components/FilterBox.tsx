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
  const [filter, setFilter] = React.useState("フォロー");

  const handleFilter = useDebouncedCallback((term: string) => {
    console.log(term);
    if (term === "follow") {
      setFilter("フォロー");
    } else {
      setFilter("フォロワー");
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
      <SelectTrigger className="w-[180px] border-2 border-primary rounded-xl text-white p-2">
        <SelectValue placeholder="フォロー" className="text-white">
          {filter}
        </SelectValue>
        <ChevronDownIcon className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400" />
      </SelectTrigger>
      <SelectContent className="bg-black text-white w-[180px]">
        <SelectGroup>
          <SelectItem
            className="text-white hover:bg-primary mb-2"
            value="follow"
            defaultChecked
          >
            フォロー
          </SelectItem>
          <SelectItem className="text-white hover:bg-primary" value="follower">
            フォロワー
          </SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};
