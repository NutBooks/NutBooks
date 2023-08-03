import Home from '@/app/(service)/page';
import { render, screen } from '@testing-library/react';

describe('<Home />', () => {
  it('renders a heading', () => {
    const { container } = render(<Home />);

    const home = screen.getByText('app/page.tsx');

    expect(home).toBeInTheDocument();
    expect(container).toMatchSnapshot();
  });
});
