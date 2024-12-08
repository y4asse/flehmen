import { Flex } from "@/components/ui/flex";
import { fetchGet } from "@/lib/fetcher";

type Props = {
  query: string;
};

export type University = {
  id: number;
  name: string;
  deviationLowerValue: number;
  deviationUpperValue: number;
  prefecture: string;
  abbreviation: string;
};

const UniversityList = async (props: Props) => {
  const { query } = props;
  const DbResult = await fetchGet<University[]>("/universities");

  const searchResult = DbResult.filter(
    (university) =>
      university.abbreviation === query ||
      university.name === query ||
      university.abbreviation.toLowerCase() == query
  );

  return (
    <Flex className="gap-4 mb-4" direction={"column"}>
      {searchResult.map((university) => (
        <Flex
          key={university.id}
          className="gap-2"
          direction={"column"}
          justify={"start"}
          align={"start"}
        >
          <p className="text-white">{university.name}</p>
          <Flex className="gap-1 pl-8">
            <p className="text-white">偏差値：0</p>
            <p className="text-white">〜</p>
            <p className="text-white">{0}</p>
            <p className="text-white ml-4">場所：{university.prefecture}</p>
            <p className="text-white ml-4">略称：{university.abbreviation}</p>
          </Flex>
        </Flex>
      ))}
    </Flex>
  );
};

export default UniversityList;
