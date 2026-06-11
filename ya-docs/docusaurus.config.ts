import {themes as prismThemes} from 'prism-react-renderer';
import type {Config} from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

// This runs in Node.js - Don't use client-side code here (browser APIs, JSX...)

const config: Config = {
  title: 'Ya',
  tagline: 'Run your commands. Right now.',
  favicon: 'img/favicon.ico',

  future: {
    v4: true,
  },

  url: 'https://d3uceY.github.io',
  baseUrl: '/Ya-CLI/',

  organizationName: 'd3uceY',
  projectName: 'Ya-CLI',

  onBrokenLinks: 'throw',

  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          editUrl: 'https://github.com/d3uceY/Ya-CLI/tree/main/ya-docs/',
          routeBasePath: 'docs',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    // Replace with your project's social card
    image: 'img/docusaurus-social-card.jpg',
    colorMode: {
      respectPrefersColorScheme: true,
    },
    navbar: {
      title: 'Ya',
      logo: {
        alt: 'Ya logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'docSidebar',
          sidebarId: 'tutorialSidebar',
          position: 'left',
          label: 'Docs',
        },
        {
          href: 'https://github.com/d3uceY/Ya-CLI',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Introduction',
              to: '/docs/intro',
            },
          ],
        },
        {
          title: 'Project',
          items: [
            { label: 'GitHub', href: 'https://github.com/d3uceY/Ya-CLI' },
            { label: 'Releases', href: 'https://github.com/d3uceY/Ya-CLI/releases' },
            { label: 'Ya GUI', href: 'https://github.com/d3uceY/Ya-GUI' },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Ya. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
