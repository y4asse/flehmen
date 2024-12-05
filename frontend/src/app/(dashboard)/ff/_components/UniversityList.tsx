import { Flex } from "@/components/ui/flex";

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
  const searchResult = universityList.filter(
    (university) =>
      university.abbreviation.includes(query) ||
      university.name.includes(query) ||
      university.abbreviation.toLowerCase().includes(query)
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
            <p className="text-white">
              偏差値：{university.deviationLowerValue}
            </p>
            <p className="text-white">〜</p>
            <p className="text-white">{university.deviationUpperValue}</p>
            <p className="text-white ml-4">場所：{university.prefecture}</p>
          </Flex>
        </Flex>
      ))}
    </Flex>
  );
};

const universityList: University[] = [
  {
    id: 1,
    name: "東京大学",
    deviationLowerValue: 60,
    deviationUpperValue: 70,
    prefecture: "東京都",
    abbreviation: "UT",
  },
  {
    id: 2,
    name: "京都大学",
    deviationLowerValue: 60,
    deviationUpperValue: 70,
    prefecture: "京都府",
    abbreviation: "KU",
  },
  {
    id: 3,
    name: "大阪大学",
    deviationLowerValue: 60,
    deviationUpperValue: 65,
    prefecture: "大阪府",
    abbreviation: "OU",
  },
];
export default UniversityList;
