type Title = string;
type Link = string;
type Bookmark_id = number;
type Keyword = string;

interface AddBookmark {
  title?: Title;
  link: Link;
}

interface Post {
  id: number;
  title: string;
  link: string;
  keywords: string[];
  status: boolean;
}
