export function Footer() {
    return (
        <footer className="flex py-10 justify-center items-center max-w-screen">
            <ul className="flex flex-row text-primary gap-24">
                <li>
                    <a href="https://hack.tum.de/" className="hover:font-medium" target={"_blank"}>About</a>
                </li>
                <li>
                    <a href="https://devpost.com/software/givemebloombergterminalplease" className="hover:font-medium" target={"_blank"}>Team</a>
                </li>
                <li>
                    <a href="https://github.com/HackaTUM-2022-PMWO/GameOfStonks" className="hover:font-medium" target={"_blank"}>GitHub</a>
                </li>
            </ul>
        </footer>
    );
}