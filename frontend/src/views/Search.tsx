import { Card } from "../components/card/Card";
import { StonkName } from "../services/vo-stonks";
import { useState } from "react";
import { StonkPositionList } from "../components/listItems/StonkPositionList";

function Search() {
  const [query, setQuery] = useState("");

  return (
    <div className="flex flex-col min-h-screen mt-5 mx-5">
      <div className="mx-10">
        <form>
          <label
            htmlFor="default-search"
            className="mb-2 text-sm font-medium text-gray-900 sr-only"
          >
            Search
          </label>
          <div className="relative">
            <div className="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
              <svg
                aria-hidden="true"
                className="w-5 h-5 text-gray-500"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                ></path>
              </svg>
            </div>
            <input
              type="search"
              id="default-search"
              value={query}
              onChange={(event) => setQuery(event.target.value)}
              className="block w-full p-4 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500"
              placeholder="Search Stonks..."
              required
            />
            <button
              type="submit"
              className="text-primary absolute right-2.5 bottom-2.5 bg-accent2 focus:ring-4 focus:outline-none hover:font-semibold focus:ring-purple-300 font-medium rounded-lg text-sm px-4 py-2 "
            >
              Search
            </button>
          </div>
        </form>
      </div>
      <Card>
        <StonkPositionList
          stonks={Object.values(StonkName)
            .filter((val) => val.valueOf().includes(query) && val !== "")
            .reduce((acc, entry) => (acc[entry] = 0), {} as any)}
        />
      </Card>
    </div>
  );
}

export default Search;
