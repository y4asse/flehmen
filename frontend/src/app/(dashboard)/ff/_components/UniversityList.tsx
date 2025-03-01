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
      university.abbreviation.toLowerCase().includes(query.toLowerCase()) ||
      university.name.toLowerCase().includes(query.toLowerCase())
  );
  return (
    <Flex className="gap-4 mb-4 px-4 " direction={"column"}>
      {searchResult.map((university) => (
        <Flex
          key={university.id}
          className="gap-2"
          direction={"column"}
          justify={"start"}
          align={"start"}
        >
          <p className="text-white">{university.name}</p>
          <Flex className="gap-1 ~pl-0/8">
            <p className="text-white">
              偏差値：{university.deviationLowerValue}
            </p>
            <p className="text-white">〜</p>
            <p className="text-white">{university.deviationUpperValue}</p>
            <p className="text-white ~ml-2/4">場所：{university.prefecture}</p>
            <p className="text-white ~ml-2/4">
              略称：{university.abbreviation}
            </p>
          </Flex>
        </Flex>
      ))}
    </Flex>
  );
};

export default UniversityList;
